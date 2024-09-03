import { ListBox } from "primereact/listbox";
import { useEffect, useState } from "react";
import { getWinningMatchup, getTeamName } from "../api/winners";
import { DataTable } from "primereact/datatable";


export default function WinningMatchup() {

    const [matchupData, setMatchupData] = useState("")
    const [selectedTeam, setSelectedTeam] = useState("")
    const [selectedTeamData, setSelectedTeamData] = useState([])
    const [teamNameData, setTeamNameData] = useState([])
    const [winningTeam, setWinningTeam] = useState([])

    useEffect(() => {
        const getMatchups = async () => {

        
        var data = await getWinningMatchup()

        var teamNames = await getTeamName()

        
        console.log(data)

        console.log(teamNames)

        setMatchupData(data)
        setTeamNameData(teamNames)

        }




        getMatchups()

        


          
    
    }, []);


    let res;
    function changeSelectedTeam(val) {

        console.log(val)
        let d = matchupData[val]
        console.log(d)

        let matchupCount = d["WinningMatchupCount"]

        let wins = d["WinningMatchupTeams"]

        let winningTeam = Object.keys(wins)


        console.log(matchupCount, wins, winningTeam)


        setWinningTeam(winningTeam)
        
    }


    let button;
    if (!teamNameData) {
        button = <ListBox options={["egg"]} />
    }
    else {
        button = <ListBox value={selectedTeam} onChange={(e) => changeSelectedTeam(e.value)} options={teamNameData} />
    }
    
    

    
        return (
            <div>
            <div>{button}</div>
            <ListBox options={winningTeam} />
            </div>
        );
    
}