import { ethers } from 'ethers';
import ICoprocessorTaskManagerABI from './abis/ICoprocessorTaskManager.json';

const CONTRACT_ADDRESS = '0xb51032037ee0FB6F2d09BE6Dbe44556bADF370A1';


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
