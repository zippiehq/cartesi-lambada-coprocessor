import React, { useState } from 'react';
import { Container, Typography } from '@mui/material';
import ConnectWallet from './components/connectWallet';
import EventListener from './components/eventListener';
import SendTask from './components/sendTask';

function App() {
    const [account, setAccount] = useState(null);

    const handleAccountChange = (account) => {
        console.log('Connected account:', account);
        setAccount(account);
    };

    return (
        <Container maxWidth="lg" style={{ marginTop: '2rem' }}>
            <Typography variant="h3" gutterBottom>
                Coprocessor Explorer
            </Typography>
            <ConnectWallet onAccountChange={handleAccountChange} />
            <SendTask account={account} />
            <EventListener />
        </Container>
    );
}

export default App;
