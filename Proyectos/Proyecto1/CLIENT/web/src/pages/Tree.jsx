import {useState} from 'react'
import { FaEye } from "react-icons/fa";
import TreeCanva from '../components/TreeCanva'
import { MdCleaningServices } from "react-icons/md";
//import data2 from './data2.json'

const Tree = () => {
  const [listProcess, setListProcess] = useState([])
  const [selectedValue, setSelectedValue] = useState('')
  const [tree, setTree] = useState(null)

  const getProcess = async () => {
    try {
      const response = await fetch('http://localhost:8080/listProcess')
      const data = await response.json()
      setListProcess(data.message.processes)
    } catch (error) {
      console.log('Error: ', error)
    }
  }

  useState(() => {
    getProcess()
  }, [])


 
  const filteredPID = (pid) => {
    const process = listProcess.filter(process => process.pid == pid)
    return process
  }

  const showValue = () => {
    if (selectedValue === '') return
      const process = filteredPID(selectedValue)
      setTree(process)
  }

  const cleanTree = () => {
    setTree(null)
  }
  
  const handleChange = (event) => {
    setSelectedValue(event.target.value);
  };


  
  return (
    <section className=''>
      <div className=' p-10'>
       <h1 className='text-4xl text-center  text-slate-700 font-bold mb-5'>Tree Process</h1>
        <button onClick={showValue} className='bg-slate-700 text-white p-2 rounded-md mr-4'>
          View
          <FaEye className='inline-block ml-2'/>
        
        </button>
        <select 
          className='w-1/4 p-2 rounded-md border-2 border-slate-700'
          value={selectedValue}
          onChange={handleChange}
        >
          <option value=''>Select a process</option>
          {
            listProcess.map((process, index) => (
              <option key={index} value={process.pid}> PID: {process.pid} - {process.name} </option>
            ))
          }
        </select>
        <button onClick={cleanTree} className='bg-orange-400 text-white p-2 rounded-md ml-4'>
          Clean
          <MdCleaningServices className='inline-block ml-2'/>
        </button>
       
       
      </div>
     
        <div className='h-96'>

          { tree && <TreeCanva data={tree} />}
        </div>
  
    </section>
  )
}

export default Tree