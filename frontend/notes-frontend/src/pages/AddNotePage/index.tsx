import AddNoteForm from "../../components/AddNoteForm"
import { useState, useEffect } from "react";
import { useNavigate } from 'react-router-dom';
import "./style.css"
import API from "../../services/API";

type NoteData = {
    title?: string,
    note: string,
    author?: string
}

export default function AddNotePage() {
    const navigate = useNavigate();

    function sendNote(data: NoteData) {
        console.log(data)
        API.post(`/notes/add`, data, {
            headers: {
                'Content-Type': 'application/json',
            }
        })
        .then(response => {
            console.log(response.data);
            navigate(`/notes/${response.data}`);
        })
        .catch(error =>{
            console.error(error);
            if(error.response.status == 404) {
                alert("Note not found");
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
        console.log(inputs);
        sendNote(inputs);
    }

    return(
    <div className="add-note-page">
        <h1>Новая заметка</h1>
        <AddNoteForm onInputChange={handleFormChange}/>
    </div>
    )
}