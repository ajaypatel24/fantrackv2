import { ListBox } from "primereact/listbox";
import { useEffect, useState } from "react";
import { getCategoryMap, getLeaderData, getTeamName } from "../api/winners";
import { DataTable } from "primereact/datatable";
import { Accordion, AccordionTab } from "primereact/accordion";
import { matchRoutes } from "react-router-dom";


export default function Leaders(props) {

    const [matchupData, setMatchupData] = useState("")
    const [selectedTeam, setSelectedTeam] = useState("")
    const [selectedCategory, setSelectedCategory] = useState([])

    useEffect(() => {
        const getMatchups = async () => {
            setMatchupData(props.data)

            console.log(props.data)
        }

        getMatchups()
    
    }, []);


    function changeSelectedTeam(val) {
        setSelectedTeam(val)
        setSelectedCategory(matchupData[val]["Value"])
    }


    let options;
    let values;

    if (!matchupData) {
        values = () => <div></div>
        options = null
    }
    else {

        values = () => { //dict of team: value
            return Object.keys(selectedCategory).map((x, i) => {
                return (
                <AccordionTab header={x}>
                    {selectedCategory[x]}
                </AccordionTab>
                );
            });
        };
        
        //array of categories
        options = <ListBox value={selectedTeam} onChange={(e) => changeSelectedTeam(e.value)} options={Object.keys(matchupData)} />
        
    }

        return (
            <div>
            <div className="grid grid-cols-12 gap-4 stick top-0">
                <div class="col-span-5">
                    {options}
                </div>

                <div class="col-span-7">

                    <Accordion>
                    {values()}
                    </Accordion>
                </div>
            
            </div>
            </div>
            
        );
    
}