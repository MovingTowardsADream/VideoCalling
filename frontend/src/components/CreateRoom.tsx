import { FormEvent } from 'react';
import {useNavigate} from "react-router-dom";

function CreateRoom() {
    const navigate = useNavigate();

    const create = async (e: FormEvent) => {
        e.preventDefault()

        const resp = await fetch("http://localhost:8080/create");
        const respJson = await resp.json();

        navigate(`/room/${respJson.roomID}`);
    }


    return (
      <div>
          <button onClick={create}>CreateRoom</button>
      </div>
    );
}

export default CreateRoom;