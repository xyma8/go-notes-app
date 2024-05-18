import "./style.css"
import { useParams } from 'react-router-dom';
import { useEffect, useState } from "react";
import { getNote } from "../../services/noteService";
import NoteContainer from "../../components/NoteContainer";

type NoteData = {
    title?: string,
    note: string,
    author?: string,
    timestamp: string
}

export default function NotePage() {
    const { noteId } = useParams();
    const [noteData, setNoteData] = useState<NoteData>();

    useEffect(() => {
        getNoteData();

    }, []);


    function getNoteData() {
        getNote(noteId)
        .then(response => {
            console.log(response.data);
            const data: NoteData = {
                title: response.data.title,
                note: response.data.note_text,
                author: response.data.author,
                timestamp: response.data.creation_time
            }

            setNoteData(data);
        })
        .catch(error => {
            console.error(error);

            if(error.response.status == 409) {
                alert("This track file already exists");
            }
            else if(error.response.status != 200) {
                alert("Internal error");
            }
        })
    }

    if(!setNoteData) {
        return(<div></div>)
    }

    return(
    <div className="note-page">
        <NoteContainer title={noteData?.title} note={noteData?.note} author={noteData?.author} timestamp={noteData?.timestamp}/>
    </div>
    )
}