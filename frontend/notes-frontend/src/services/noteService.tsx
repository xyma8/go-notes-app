import API from './API';

export function getNote(noteId: string | undefined) {
    return API.get(`/notes/${noteId}`)
}