import { useEffect, useState } from 'react'
import './App.css'
import { Canvas } from './components/canvas/Canvas'
import { readShapesFromFile } from './utils/readShapes'
import { Shape } from './types/shape/shape'
import { exportPerimeterAndArea } from './utils/writeShapes'

function App() {
  const [shapes, setShapes] = useState<Shape[]>([])

  useEffect(() => {
    exportPerimeterAndArea(shapes);
  }, [shapes]);

  return (
    <>
      <input type="file" onChange={async e =>
            e.target.files && setShapes(await readShapesFromFile(e.target.files[0]))} />
      <Canvas shapes={shapes} />
    </>
  )
}

export default App
