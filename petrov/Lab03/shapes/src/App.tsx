import { useEffect, useRef, useState } from 'react'
import './App.css'
import { Canvas } from './components/canvas/Canvas'
import { readShapesFromFile } from './utils/readShapes'
import { Shape } from './types/shape/shape'
import { PerimeterAndAreaVisitor } from './types/visitor/perimeterAndAreaVisitor'

function App() {
  const [shapes, setShapes] = useState<Shape[]>([]);
  const perimeterAndAreaVisitorRef = useRef<PerimeterAndAreaVisitor>(new PerimeterAndAreaVisitor());

  useEffect(() => {
    // exportPerimeterAndArea(shapes);
    shapes.forEach(shape => shape.accept(perimeterAndAreaVisitorRef.current));
    perimeterAndAreaVisitorRef.current.downloadData('output.txt');
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
