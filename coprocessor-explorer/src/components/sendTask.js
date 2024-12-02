import React, { useState } from 'react';
import { Button, Typography, Paper, Box, Alert, TextField } from '@mui/material';

const SendTask = ({ account }) => {
    const [programId, setProgramId] = useState('');
    const [file, setFile] = useState(null);
    const [loading, setLoading] = useState(false);
    const [successMessage, setSuccessMessage] = useState('');
    const [errorMessage, setErrorMessage] = useState('');

    const handleFileChange = (event) => {
        setFile(event.target.files[0]);
    };

    const handleSubmit = async (e) => {
        e.preventDefault();

        setSuccessMessage('');
        setErrorMessage('');

        if (!account) {
            setErrorMessage('Please connect your wallet first.');
            return;
        }

        if (!programId || !file) {
            setErrorMessage('Both Program ID and a file are required.');
            return;
        }

        setLoading(true);

        try {
            const reader = new FileReader();
            reader.readAsDataURL(file);
            reader.onload = async () => {
                const base64Input = reader.result.split(',')[1];

                const payload = JSON.stringify([
                    {
                        programId: programId,
                        input: base64Input,
                    },
                ]);

                const response = await fetch('http://lambada.tspre.org:8090/createTask', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: payload,
                });

                const data = await response.json();
                setSuccessMessage('Task sent successfully!');
                setProgramId('');
                setFile(null);
                console.log('Response:', data);
            };
            reader.onerror = (error) => {
                throw error;
            };
        } catch (error) {
            console.error('Error:', error);
            setErrorMessage(error.message || 'An unexpected error occurred.');
        } finally {
            setLoading(false);
        }
    };

    return (
        <Paper elevation={3} style={{ padding: '2rem', marginBottom: '2rem' }}>
            <Typography variant="h5" gutterBottom>
                Send a New Task
            </Typography>
            <form onSubmit={handleSubmit}>
                <Box display="flex" flexDirection="column" gap="1rem">
                    <TextField
                        label="Program ID (CID)"
                        variant="outlined"
                        value={programId}
                        onChange={(e) => setProgramId(e.target.value)}
                        required
                        fullWidth
                    />
                    <input
                        type="file"
                        onChange={handleFileChange}
                        required
                    />
                    <Button
                        type="submit"
                        variant="contained"
                        color="primary"
                        disabled={loading || !file || !account} // Disable button if wallet is not connected
                    >
                        {loading ? 'Sending...' : 'Send Task'}
                    </Button>
                    {!account && (
                        <Typography color="error">Please connect your wallet to send a task.</Typography>
                    )}
                </Box>
            </form>
            {successMessage && (
                <Box mt={2}>
                    <Alert severity="success">{successMessage}</Alert>
                </Box>
            )}
            {errorMessage && (
                <Box mt={2}>
                    <Alert severity="error">{errorMessage}</Alert>
                </Box>
            )}
        </Paper>
    );
};

export default SendTask;
