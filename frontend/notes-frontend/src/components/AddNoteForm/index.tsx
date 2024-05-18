import "./style.css"

export default function AddNoteForm() {

    return(
    <div className="add-note-form">
        <input autoComplete="off" placeholder="Тема заметки (необязательно)"></input>
        <textarea autoComplete="off"
         placeholder="Заметка" 
         rows={20} 
         cols={80}>
        </textarea>
        <input autoComplete="off" placeholder="Author" defaultValue={"Аноним"}></input>
    </div>
    )
}