import { ethers } from 'ethers';
import ICoprocessorTaskManagerABI from './abis/ICoprocessorTaskManager.json';

const CONTRACT_ADDRESS = '0x000000000000000000000000000000000000dEaD';

if (CONTRACT_ADDRESS === '0x000000000000000000000000000000000000dEaD') {
    console.warn('Using dummy contract address.');
}

export const getProvider = () => {
    if (window.ethereum) {
        return new ethers.providers.Web3Provider(window.ethereum);
    } else {
        console.error('Please install MetaMask!');
        return null;
    }
};

export const getContract = (provider) => {
    return new ethers.Contract(CONTRACT_ADDRESS, ICoprocessorTaskManagerABI, provider);
};
