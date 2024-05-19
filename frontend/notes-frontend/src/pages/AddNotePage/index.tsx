import AddNoteForm from "../../components/AddNoteForm"
import { useState, useEffect } from "react";
import "./style.css"
import API from "../../services/API";

type NoteData = {
    title?: string,
    note: string,
    author?: string
}

export default function AddNotePage() {
    const [noteData, setNoteData] = useState<NoteData>();

    useEffect(() => {
        sendNote();

    }, [noteData]);

    function sendNote() {
        API.post(`/notes/add`, noteData)
        .then(response => {
            console.log(response.data);
        })
        .catch(error =>{
            console.error(error);
            if(error.response.status == 404) {
                alert("Workout not found");
            }
            if(error.response.status == 403) {
                alert("Access denied");
            }
            else if(error.response.status != 200) {
                alert("Internal error");
            }
        })
    }

    function handleFormChange(inputs: NoteData) {
        setNoteData(inputs);
        //sendNote(inputs);
    }

    return(
    <div className="add-note-page">
        <h1>Новая заметка</h1>
        <AddNoteForm onInputChange={handleFormChange}/>
    </div>
    )
}