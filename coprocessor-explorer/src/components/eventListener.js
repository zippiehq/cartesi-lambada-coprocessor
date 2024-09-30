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
        const provider = getProvider();
        if (!provider) return;

        const contract = getContract(provider);

        // TaskResponded events
        contract.on('TaskResponded', (task, response) => {
            let decodedProgramId;
            let decodedResultCID;

            try {
                decodedProgramId = CID.decode(Buffer.from(task.programId.slice(2), 'hex')).toString();
                decodedResultCID = CID.decode(Buffer.from(response.resultCID.slice(2), 'hex')).toString();
            } catch (error) {
                console.error('Error decoding CID:', error);
                decodedProgramId = 'Invalid CID';
                decodedResultCID = 'Invalid CID';
            }

            const newResponse = {
                batchIndex: task.batchIndex.toNumber(),
                programId: decodedProgramId,
                inputHash: ethers.utils.hexlify(task.inputHash),
                resultCID: decodedResultCID,
                outputHash: ethers.utils.hexlify(response.outputHash),
            };
            setTaskResponses((prev) => [newResponse, ...prev]);
            console.log('TaskResponded:', newResponse);
        });

        // TaskBatchRegistered events
        contract.on('TaskBatchRegistered', (batch) => {
            const newBatch = {
                index: batch.index.toNumber(),
                blockNumber: batch.blockNumber.toNumber(),
                merkeRoot: ethers.utils.hexlify(batch.merkeRoot),
                quorumNumbers: ethers.utils.hexlify(batch.quorumNumbers),
                quorumThresholdPercentage: batch.quorumThresholdPercentage.toNumber(),
            };
            setTaskBatches((prev) => [newBatch, ...prev]);
            console.log('TaskBatchRegistered:', newBatch);
        });

        // Cleanup
        return () => {
            contract.removeAllListeners('TaskResponded');
            contract.removeAllListeners('TaskBatchRegistered');
        };
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
