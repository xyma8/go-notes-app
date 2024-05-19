import "./style.css"
import { useState } from "react";

type NoteData = {
    title?: string,
    note: string,
    author?: string
}

type AddNoteFormProps = {
    onInputChange: (inputs: NoteData) => void
}

export default function AddNoteForm(props: AddNoteFormProps) {
    const [formData, setFormData] = useState<NoteData>({ title: "", note: "", author: "Аноним" })

    function handleTitleChange(event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) {
        const value = event.target.value;
        setFormData(prevState => ({ ...prevState, title: value }));
    };

    function handleNoteChange(event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) {
        const value = event.target.value;
        setFormData(prevState => ({ ...prevState, note: value }));
    };

    function handleAuthorChange(event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) {
        const value = event.target.value;
        setFormData(prevState => ({ ...prevState, author: value }));
    };

    function sendNote() {
        if(formData) {
            props.onInputChange(formData);
        }
    }

    return(
    <div className="add-note-form">
        <input autoComplete="off" placeholder="Заголовок заметки (необязательно)" onChange={handleTitleChange}></input>
        <textarea autoComplete="off"
         placeholder="Заметка" 
         rows={20} 
         cols={80}
         onChange={handleNoteChange}>
        </textarea>
        <input autoComplete="off" placeholder="Author" defaultValue={"Аноним"} onChange={handleAuthorChange}></input>
        <button onClick={sendNote} className="save-button">Сохранить</button>
    </div>
    )
}