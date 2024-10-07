import React, { useEffect, useState, useRef } from 'react';
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
    Link,
} from '@mui/material';
import { getProvider, getContract } from '../config';

const EventListener = () => {
    const [taskResponses, setTaskResponses] = useState([]);
    const [taskBatches, setTaskBatches] = useState([]); // For TaskBatchRegistered events
    const [page, setPage] = useState(0);
    const [rowsPerPage, setRowsPerPage] = useState(10);

    const listenersAddedRef = useRef(false);

    useEffect(() => {
        console.log('EventListener component mounted');

        const setupListeners = async () => {
            try {
                const provider = getProvider();
                if (!provider) {
                    console.error('Provider not found');
                    return;
                }
                console.log('Provider connected:', provider);

                const contract = getContract(provider);
                if (!contract) {
                    console.error('Contract not found');
                    return;
                }

                console.log('Contract address:', contract.address);

                // Ensure we're connected to the correct network
                const network = await provider.getNetwork();
                console.log('Connected to network:', network.name);

                await fetchPastEvents(contract);

                // Add event listeners if not already added
                if (!listenersAddedRef.current) {
                    console.log("adding listeners");

                    // Listener for TaskResponded
                    contract.on('TaskResponded', (task, response) => {
                        console.log('Received TaskResponded event:', { task, response });
                        const newResponse = parseTaskResponded(task, response);
                        if (newResponse) {
                            setTaskResponses((prev) => [newResponse, ...prev]);
                            console.log('New TaskResponded event added to state');
                        }
                    });

                    // Listener for TaskBatchRegistered
                    contract.on('TaskBatchRegistered', (batch, event) => {
                        console.log('Received TaskBatchRegistered event:', batch);
                        const { index, blockNumber, merkeRoot, quorumNumbers, quorumThresholdPercentage } = batch;
                        const newBatch = parseTaskBatchRegistered(index, blockNumber, merkeRoot, quorumNumbers, quorumThresholdPercentage);
                        console.log({newBatch});
                        if (newBatch) {
                            setTaskBatches((prev) => [newBatch, ...prev]);
                            console.log('New TaskBatchRegistered event added to state');
                        }
                    });

                    listenersAddedRef.current = true;
                    console.log('Event listeners added');
                }
            } catch (error) {
                console.error('Error setting up listeners:', error);
            }
        };

        const fetchPastEvents = async (contract) => {
            try {
                console.log('Fetching past events');

                // past TaskResponded events
                const taskRespondedEvents = await contract.queryFilter('TaskResponded');
                const pastTaskResponses = taskRespondedEvents.map((event) => {
                    const { task, response } = event.args;
                    return parseTaskResponded(task, response);
                });
                console.log('past task responses fetched:', pastTaskResponses);

                setTaskResponses(pastTaskResponses.reverse());

                // past TaskBatchRegistered events
                const taskBatchRegisteredEvents = await contract.queryFilter('TaskBatchRegistered');
                const pastTaskBatches = taskBatchRegisteredEvents.map((event) => {
                    const batch = event.args.batch;
                    const { index, blockNumber, merkeRoot, quorumNumbers, quorumThresholdPercentage } = batch;

                    return parseTaskBatchRegistered(index, blockNumber, merkeRoot, quorumNumbers, quorumThresholdPercentage);
                }).filter((batch) => batch !== null);

                console.log('past task batches fetched:', pastTaskBatches);

                setTaskBatches(pastTaskBatches.reverse());

                console.log('Past events fetched');
            } catch (error) {
                console.error('Error fetching past events:', error);
            }
        };

        setupListeners();

        // Cleanup on unmount
        return () => {
            const provider = getProvider();
            const contract = getContract(provider);
            if (contract) {
                contract.removeAllListeners('TaskResponded');
                contract.removeAllListeners('TaskBatchRegistered');
                console.log('Event listeners removed');
            }
        };
    }, []);

    const parseTaskResponded = (task, response) => {
        const { batchIndex, programId, inputHash } = task;
        const { resultCID, outputHash } = response;

        let decodedProgramId = 'Invalid CID';
        let decodedResultCID = 'Invalid CID';

        try {
            const programIdBytes = ethers.utils.arrayify(programId);
            decodedProgramId = new TextDecoder('utf-8').decode(programIdBytes);

            let resultCIDBytes = resultCID;
            if (typeof resultCID === 'string') {
                resultCIDBytes = ethers.utils.arrayify(resultCID);
            }
            decodedResultCID = CID.decode(resultCIDBytes).toString();
        } catch (error) {
            console.error('Error decoding CIDs:', error);
        }

        return {
            batchIndex,
            programId: decodedProgramId,
            inputHash: ethers.utils.hexlify(inputHash),
            resultCID: decodedResultCID,
            outputHash: ethers.utils.hexlify(outputHash),
        };
    };

    const parseTaskBatchRegistered = (index, blockNumber, merkeRoot, quorumNumbers, quorumThresholdPercentage) => {
        if (merkeRoot === undefined || quorumNumbers === undefined) {
            console.warn('undefined taskBatchRegistered event');
            return null;
        }

        return {
            index,
            blockNumber,
            merkeRoot: ethers.utils.hexlify(merkeRoot),
            quorumNumbers: ethers.utils.hexlify(quorumNumbers),
            quorumThresholdPercentage,
        };
    };

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
                                    <TableCell>
                                        {response.programId !== 'Invalid CID' ? (
                                            <Link
                                                href={`http://espresso.tspre.org:8080/ipfs/${response.programId}`}
                                                target="_blank"
                                                rel="noopener noreferrer"
                                            >
                                                {response.programId}
                                            </Link>
                                        ) : (
                                            response.programId
                                        )}
                                    </TableCell>
                                    <TableCell>{response.inputHash}</TableCell>
                                    <TableCell>
                                        {response.resultCID !== 'Invalid CID' ? (
                                            <Link
                                                href={`http://espresso.tspre.org:8080/ipfs/${response.resultCID}`}
                                                target="_blank"
                                                rel="noopener noreferrer"
                                            >
                                                {response.resultCID}
                                            </Link>
                                        ) : (
                                            response.resultCID
                                        )}
                                    </TableCell>
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
                            .filter(batch => batch !== null)
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
