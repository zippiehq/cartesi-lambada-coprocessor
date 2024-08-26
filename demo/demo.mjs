import ethers from "ethers";
import fs from "fs";
import crypto from "crypto"
import { CID } from "multiformats";
import { fileURLToPath } from 'url';
import { dirname } from 'path';

const providerUrl = "http://localhost:8545";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);
const deploymentPath = __dirname + "/../contracts/script/output/coprocessor_deployment_output_devnet.json";
const abiFilePath = __dirname + "/../contracts/out/ICoprocessorTaskManager.sol/ICoprocessorTaskManager.json"; 

const contractAddress = JSON.parse(fs.readFileSync(deploymentPath, 'utf8')).addresses.taskManager;

// Replace with actual programId and taskInputHash
const programId = process.env["PROGRAM_CID"] ? process.env["PROGRAM_CID"] : "bafybeicdhhtwmgpnt7jugvlv3xtp2u4w4mkunpmg6txkkkjhpvnt2buyqa";
const taskInput = process.env["PROGRAM_INPUT"] ? process.env["PROGRAM_INPUT"] : "echo hello world";
const taskInputBuf = Buffer.from(taskInput, "utf8");
const taskInputHash = ethers.utils.keccak256(taskInputBuf);
console.log("taskInputHash: " + taskInputHash);

// Load the ABI from the JSON file
const contractABI = JSON.parse(fs.readFileSync(abiFilePath, 'utf8')).abi;

// Connect to the Ethereum provider
const provider = new ethers.providers.JsonRpcProvider(providerUrl);

// Create a contract instance
const contract = new ethers.Contract(contractAddress, contractABI, provider);

// Data structure to map block numbers and hashes to batch indices
let taskBatchRegistry = {};

let nextBatchIndex = -1;

// Function to initialize and fetch the next batch index
async function initialize() {
    nextBatchIndex = await contract.getNextBatchIndex();
    console.log("Next Batch Index:", nextBatchIndex.toString());
}

// Function to handle TaskBatchRegistered event
function handleTaskBatchRegistered(index, blockNumber, merkeRoot, quorumNumbers, quorumThresholdPercentage) {
    taskBatchRegistry[index] = { blockNumber, merkeRoot, quorumNumbers, quorumThresholdPercentage };
    console.log("TaskBatchRegistered:", index, blockNumber, merkeRoot);
}

// Function to handle TaskResponded event
async function handleTaskResponded(task, response) {
    console.log(response);
    const { batchIndex, programId: respProgramId, inputHash: respTaskInputHash } = task;
    const { resultCID, outputHash } = response;
    if (batchIndex >= nextBatchIndex && Buffer.from(respProgramId.slice(2), "hex").toString("utf8") === programId && respTaskInputHash === taskInputHash) {
        console.log("Task Responded:", {
            batchIndex,
            outputHash,
            resultCID: response.resultCID
        });

        const blockNumber = taskBatchRegistry[batchIndex]?.blockNumber;
        const blockHash = taskBatchRegistry[batchIndex]?.merkeRoot;

        console.log("Ethereum Block Number:", blockNumber);
        console.log("Ethereum Block Hash:", blockHash);
        
        const blockNumberBuf = Buffer.alloc(8);
        blockNumberBuf.writeBigUint64BE(BigInt(blockNumber));
        const blockHashBuf = Buffer.from(blockHash.slice(2), "hex");
        let computeJob = {
            metadata: {},
            payload: taskInputBuf.toString("base64"),
        }
        computeJob.metadata[crypto.createHash("sha256").update(Buffer.from("sequencer", "utf8")).digest("base64")] =
            crypto.createHash("sha256").update(Buffer.from("coprocessor", "utf8")).digest("base64");
        
        computeJob.metadata[crypto.createHash("sha256").update(Buffer.from("coprocessor-batch-block-number", "utf8")).digest("base64")] = 
            blockNumberBuf.toString("base64");
        computeJob.metadata[crypto.createHash("sha256").update(Buffer.from("coprocessor-batch-block-hash", "utf8")).digest("base64")] =
            blockHashBuf.toString("base64");
            

        console.log(computeJob);
        let newResultCID = (await computeResult(programId, computeJob)).cid;
        let output = await ipfsGet(newResultCID + "/output");
        let cid = CID.decode(Buffer.from(resultCID.slice(2), "hex"));
        let r2 = crypto.createHash("sha256").update(output).digest();
        if (outputHash.slice(2) === r2.toString("hex") && cid.toString() == newResultCID) {
            console.log("Local compute job matches output of co-processor, result CID " + cid.toString() + " output file hash: " + r2.toString("hex"));
            console.log("output file contents: " + output.toString("utf8"));
            process.exit(0);
        }
    }
}

// Event listeners
contract.on("TaskBatchRegistered", (batch) => {
    handleTaskBatchRegistered(batch.index, batch.blockNumber, batch.merkeRoot, batch.quorumNumbers, batch.quorumThresholdPercentage);
});

contract.on("TaskResponded", (task, response) => {
    handleTaskResponded(task, response);
});

async function sendTaskInput(programId, taskInput) {
    const url = 'http://127.0.0.1:8080/createTask';
    const data = [
        {
            "programId": programId,
            "input": Buffer.from(taskInput, "utf8").toString("base64")
        }
    ];

    try {
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });

        if (!response.ok) {
            throw new Error('Network response was not ok ' + response.statusText);
        }

        const responseData = await response.json();
        return responseData;
    } catch (error) {
        console.error('There was a problem with the fetch operation:', error);
    }
}

async function computeResult(programId, computeJob) {
    const url = 'http://127.0.0.1:3001/compute/' + programId + '?json=true';

    try {
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/octet-stream'
            },
            body: JSON.stringify(computeJob)
        });

        if (!response.ok) {
            throw new Error('Network response was not ok ' + response.statusText);
        }

        const responseData = await response.json();
        console.log(responseData);
        return responseData;
    } catch (error) {
        console.error('There was a problem with the fetch operation:', error);
    }
}

async function ipfsGet(ipfsPath) {
    const url = 'http://127.0.0.1:5001/api/v0/cat?arg=' + ipfsPath;

    try {
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/octet-stream'
            },
            body: '',
        });

        if (!response.ok) {
            throw new Error('Network response was not ok ' + response.statusText);
        }

        const responseData = await response.arrayBuffer();
        return Buffer.from(responseData);
    } catch (error) {
        console.error('There was a problem with the fetch operation:', error);
    }
}

async function init() {
    await initialize();
    await sendTaskInput(programId, taskInput);
}

init().catch((err) => { console.log(err); });