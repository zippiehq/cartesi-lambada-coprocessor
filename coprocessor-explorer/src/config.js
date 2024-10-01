import { ethers } from 'ethers';
import ICoprocessorTaskManagerABI from './abis/ICoprocessorTaskManager.json';

const CONTRACT_ADDRESS = '0xf1a22A655aA923dF47893a44F3765Bf30A7DD0a8';


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
