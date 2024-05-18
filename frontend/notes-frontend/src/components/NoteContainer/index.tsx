import "./style.css"

type NoteContainerProps = {
    title?: string | undefined,
    note: string | undefined
    author?: string | undefined
    timestamp: string | undefined
}

export default function NoteContainer(props: NoteContainerProps) {

    return(
    <div className="note-container">
        {props?.title && <h2 className="note-display-title">{props.title}</h2>}
        <p className="note-display-content">{props.note}</p>
        <div className="note-display-footer">
        <span className="note-display-author">Автор: {props.author}</span>
        {props?.timestamp && <span className="note-display__time">Создано: {new Date(props.timestamp).toLocaleString()}</span>}
      </div>
    </div>
    )
}