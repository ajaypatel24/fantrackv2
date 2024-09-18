import apiClient from "./client";
import axios from 'axios'


export const getWinners = async () => {
    const response = await axios.get(`http://localhost:8080/login`
       
    );

    return response.data
}

export const getLeaderData = async () => {
    const response = await axios.get(`http://localhost:8080/winners`, 
        { withCredentials: true}
    );

    return response.data
}

export const getCategoryMap = async () => {
    const response = await axios.get(`http://localhost:8080/category`,
        { withCredentials: true}
    );

    return response.data
}


export const getWinningMatchup = async () => {
    const response = await axios.get(`http://localhost:8080/matchups`,
        {withCredentials: true}
    )

    return response.data
}


export const getTeamName = async () => {
    const response = await axios.get(`http://localhost:8080/teams`,
        {withCredentials: true}
    )

    return response.data
}