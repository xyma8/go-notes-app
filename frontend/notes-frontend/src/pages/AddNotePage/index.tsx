import AddNoteForm from "../../components/AddNoteForm"
import "./style.css"

export default function AddNotePage() {

    function sendNote() {

    }

    return(
    <div className="add-note-page">
        <h1>Новая заметка</h1>
        <AddNoteForm/>
        <button onClick={sendNote} className="save-button">Сохранить</button>
    </div>
    )
}