import React, { useEffect, useState } from 'react';
import { ethers } from 'ethers';
import { CID } from 'multiformats/cid';
import {
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
    Paper,
    Typography,
    TablePagination,
} from '@mui/material';
import { getProvider, getContract } from '../config';

const EventListener = () => {
    const [taskResponses, setTaskResponses] = useState([]);
    const [taskBatches, setTaskBatches] = useState([]); // For TaskBatchRegistered events
    const [page, setPage] = useState(0);
    const [rowsPerPage, setRowsPerPage] = useState(10);

    useEffect(() => {
        const setupListeners = async () => {
            const provider = getProvider();
            if (!provider) return;
            console.log('Provider:', provider.connection.url);

            const contract = getContract(provider);

            // Log contract address
            console.log('Contract address:', contract.address);

            // Log network info
            const network = await contract.provider.getNetwork();
            console.log('Connected to network:', network);

            // Listen to TaskResponded events
            contract.on('TaskResponded', (task, response) => {
                const { batchIndex, programId, inputHash } = task;
                const { resultCID, outputHash } = response;

                let decodedProgramId = 'Invalid CID';
                let decodedResultCID = 'Invalid CID';

                try {
                    let programIdBytes = programId;
                    let resultCIDBytes = resultCID;

                    if (typeof programId === 'string') {
                        programIdBytes = ethers.utils.arrayify(programId);
                    }

                    if (typeof resultCID === 'string') {
                        resultCIDBytes = ethers.utils.arrayify(resultCID);
                    }
                    decodedProgramId = CID.decode(programIdBytes).toString();
                    decodedResultCID = CID.decode(resultCIDBytes).toString();
                } catch (error) {
                    console.error('Error decoding CID:', error);
                }

                const newResponse = {
                    batchIndex: batchIndex,
                    programId: decodedProgramId,
                    inputHash: ethers.utils.hexlify(inputHash),
                    resultCID: decodedResultCID,
                    outputHash: ethers.utils.hexlify(outputHash),
                };
                setTaskResponses((prev) => [newResponse, ...prev]);
                console.log('TaskResponded:', newResponse);
            });

            contract.on('TaskBatchRegistered', (batch) => {
                console.log('TaskBatchRegistered event received:', batch);

                const {
                    index,
                    blockNumber,
                    merkeRoot,
                    quorumNumbers,
                    quorumThresholdPercentage,
                } = batch;

                if (merkeRoot === undefined || quorumNumbers === undefined) {
                    console.error('Received undefined values in TaskBatchRegistered event.');
                    return;
                }

                const newBatch = {
                    index: index,
                    blockNumber: blockNumber,
                    merkeRoot: ethers.utils.hexlify(merkeRoot),
                    quorumNumbers: ethers.utils.hexlify(quorumNumbers),
                    quorumThresholdPercentage: quorumThresholdPercentage,
                };
                setTaskBatches((prev) => [newBatch, ...prev]);
                console.log('TaskBatchRegistered:', newBatch);
            });

            // Cleanup
            return () => {
                contract.removeAllListeners('TaskResponded');
                contract.removeAllListeners('TaskBatchRegistered');
            };
        };

        setupListeners(); // Call the async function
    }, []);

    const handleChangePage = (event, newPage) => {
        setPage(newPage);
    };

    const handleChangeRowsPerPage = (event) => {
        setRowsPerPage(parseInt(event.target.value, 10));
        setPage(0);
    };

    return (
        <div>
            <Typography variant="h5" gutterBottom>
                Task Responded Events
            </Typography>
            <TableContainer component={Paper}>
                <Table>
                    <TableHead>
                        <TableRow>
                            <TableCell>Batch Index</TableCell>
                            <TableCell>Program ID (CID)</TableCell>
                            <TableCell>Input Hash</TableCell>
                            <TableCell>Result CID</TableCell>
                            <TableCell>Output Hash</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {taskResponses
                            .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
                            .map((response, idx) => (
                                <TableRow key={idx}>
                                    <TableCell>{response.batchIndex}</TableCell>
                                    <TableCell>{response.programId}</TableCell>
                                    <TableCell>{response.inputHash}</TableCell>
                                    <TableCell>{response.resultCID}</TableCell>
                                    <TableCell>{response.outputHash}</TableCell>
                                </TableRow>
                            ))}
                    </TableBody>
                </Table>
                <TablePagination
                    component="div"
                    count={taskResponses.length}
                    page={page}
                    onPageChange={handleChangePage}
                    rowsPerPage={rowsPerPage}
                    onRowsPerPageChange={handleChangeRowsPerPage}
                    rowsPerPageOptions={[5, 10, 25]}
                />
            </TableContainer>

            <Typography variant="h5" gutterBottom style={{ marginTop: '20px' }}>
                Task Batch Registered Events
            </Typography>
            <TableContainer component={Paper}>
                <Table>
                    <TableHead>
                        <TableRow>
                            <TableCell>Batch Index</TableCell>
                            <TableCell>Block Number</TableCell>
                            <TableCell>Merke Root</TableCell>
                            <TableCell>Quorum Numbers</TableCell>
                            <TableCell>Quorum Threshold Percentage</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {taskBatches
                            .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
                            .map((batch, idx) => (
                                <TableRow key={idx}>
                                    <TableCell>{batch.index}</TableCell>
                                    <TableCell>{batch.blockNumber}</TableCell>
                                    <TableCell>{batch.merkeRoot}</TableCell>
                                    <TableCell>{batch.quorumNumbers}</TableCell>
                                    <TableCell>{batch.quorumThresholdPercentage}</TableCell>
                                </TableRow>
                            ))}
                    </TableBody>
                </Table>
                <TablePagination
                    component="div"
                    count={taskBatches.length}
                    page={page}
                    onPageChange={handleChangePage}
                    rowsPerPage={rowsPerPage}
                    onRowsPerPageChange={handleChangeRowsPerPage}
                    rowsPerPageOptions={[5, 10, 25]}
                />
            </TableContainer>
        </div>
    );
};

export default EventListener;
