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
    const [formData, setFormData] = useState<NoteData>()

    const handleInputChange = (event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        const { name, value } = event.target;
        setFormData(prevState => ({
            ...prevState!,
            [name]: value
        }));
      };

    function sendNote() {
        if(formData) {
            props.onInputChange(formData);
        }
    }

    return(
    <div className="add-note-form">
        <input autoComplete="off" placeholder="Тема заметки (необязательно)" onChange={handleInputChange}></input>
        <textarea autoComplete="off"
         placeholder="Заметка" 
         rows={20} 
         cols={80}
         onChange={handleInputChange}>
        </textarea>
        <input autoComplete="off" placeholder="Author" defaultValue={"Аноним"} onChange={handleInputChange}></input>
        <button onClick={sendNote} className="save-button">Сохранить</button>
    </div>
    )
}