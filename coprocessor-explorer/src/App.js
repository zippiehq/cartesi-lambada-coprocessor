import React from 'react';
import { Container, Typography } from '@mui/material';
import ConnectWallet from './components/connectWallet';
import EventListener from './components/eventListener';
import SendTask from './components/sendTask';

function App() {
    const handleAccountChange = (account) => {
        console.log('Connected account:', account);
    };

    return (
        <Container maxWidth="lg" style={{ marginTop: '2rem' }}>
            <Typography variant="h3" gutterBottom>
                Coprocessor Explorer
            </Typography>
            <ConnectWallet onAccountChange={handleAccountChange} />
            <SendTask />
            <EventListener />
        </Container>
    );
}

export default App;
