import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <h1>Hello World!</h1>
      <button onClick={async ()=> {
        const response =await fetch('/users')
        const data= await response.json()
        console.log(data)}}>Obtener datos</button>
        </>
  )
}

export default App
