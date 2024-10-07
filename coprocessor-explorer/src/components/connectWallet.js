import React, { useState } from 'react';
import { Button, Typography } from '@mui/material';

const ConnectWallet = ({ onAccountChange }) => {
    const [account, setAccount] = useState(null);

    const connectWallet = async () => {
        if (window.ethereum) {
            try {
                const accounts = await window.ethereum.request({ method: 'eth_requestAccounts' });
                setAccount(accounts[0]);
                onAccountChange(accounts[0]);

                // Listen for account changes
                window.ethereum.on('accountsChanged', (accounts) => {
                    setAccount(accounts[0]);
                    onAccountChange(accounts[0]);
                });
            } catch (error) {
                console.error('User rejected the request.');
            }
        } else {
            alert('Please install MetaMask!');
        }
    };

    return (
        <div>
            {account ? (
                <Typography variant="h6">Connected: {account}</Typography>
            ) : (
                <Button variant="contained" color="primary" onClick={connectWallet}>
                    Connect Wallet
                </Button>
            )}
        </div>
    );
};

export default ConnectWallet;
