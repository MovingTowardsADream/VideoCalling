import {BrowserRouter, Route, Routes} from "react-router-dom";

import CreateRoom from "./components/CreateRoom.tsx";
import Room from "./components/Room.tsx";

function App() {
  return(
      <div className="App">
        <BrowserRouter>
          <Routes>
              <Route path="/room/:roomID" element={<Room />} />
              <Route path="/" element={<CreateRoom />} />
          </Routes>
        </BrowserRouter>
      </div>
  )
}

export default App
