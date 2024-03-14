import { useState } from 'react';
import { FaSkull, FaStopCircle,FaPlay } from "react-icons/fa";
import { BiSolidMessageSquareAdd } from "react-icons/bi";

import StateCanva from '../components/StateCanva';


const SManagment = () => {
  const [initialNodes, setInitialNodes] = useState([])
  const [initialEdges, setEdges] = useState([])
  const [actualState, setActualState] = useState(null)
  const [getPid, setPid] = useState(null)




 


  function getRandomInt(min, max) {
    const minCeiled = Math.ceil(min);
    const maxFloored = Math.floor(max);
    return Math.floor(Math.random() * (maxFloored - minCeiled) + minCeiled); // The maximum is exclusive and the minimum is inclusive
  }

  const getLastNodePosition = () => {
    const value =  initialNodes[initialNodes.length - 1]
    const positionY  = value.position.y
    const positionX  = value.position.x
    return [positionX, positionY]
  }

  function cleanData() {
    setInitialNodes([])
    setEdges([])
    setAxisX(100)
    setAxisY(100)
    setPid(null)
  }

  const handleNew =  async () => {
    if (initialNodes.length > 0) cleanData()

    try {
      const response = await fetch('http://localhost:8080/createProcess')
      const data = await response.json()
      setPid(data.message)
    } catch (error) {
      console.log(error)
    } finally {
      setActualState('Running')
      const newNodeIds = ['New', 'Ready', 'Running'].map((label,index) => {  
          const id = getRandomInt(0,500).toString();
          
          return {
            id,
            position: { x: 200*index, y: 100 }, // Initial position
            data: { label },
            style: { backgroundColor: index === 2 ? '#95f5a4' : 'white' }
          
          };
        }
      );

      setInitialNodes((prevNodes) => {
        // Update style of previous nodes
        const updatedNodes = prevNodes.map((node) => ({
          ...node,
          style: { backgroundColor: 'white' } // Change to white color
        }));
      
        // Append new nodes to the updated nodes
        return [...updatedNodes, ...newNodeIds];
      });

      const newEdges = newNodeIds.slice(0, -1).map((node, index) => ({
        id: `e${node.id}-${newNodeIds[index + 1].id}`,
        source: node.id.toString(),
        target: (newNodeIds[index + 1].id),
        
      })
        
      );
      setEdges((prevEdges) => [...prevEdges, ...newEdges]);

    
    }
      

      
  };

  const handleKill = () => {
    if (initialNodes.length === 0) return
    try {
      const response = fetch(`http://localhost:8080/terminateProcess?pid=${getPid}`)
      const data = response.json()
      console.log(data.message)
    } catch (error) {
      console.log(error)
    } finally {

      setActualState('Terminated')
      let positionX  = getLastNodePosition()[0]+200
      let positionY  = getLastNodePosition()[1]
      if ( positionX > 1000 ) {
        positionX = 0
        positionY = positionY + 100
      }

      const newNodeIds = ['Terminated'].map((label) => {
        const id = getRandomInt(0,500).toString();
        return {
          id,
          position: { x: positionX , y: positionY }, // Initial position
          data: { label },
          style: { backgroundColor: '#f59595' }
        };
      });

      setInitialNodes((prevNodes) => {
        // Update style of previous nodes
        const updatedNodes = prevNodes.map((node) => ({
          ...node,
          style: { backgroundColor: 'white' } // Change to white color
        }));
      
        // Append new nodes to the updated nodes
        return [...updatedNodes, ...newNodeIds];
      });

      const getLastNode = initialNodes[initialNodes.length - 1]
      const newEdges = [{
        id: `e${getLastNode.id}-${newNodeIds[0].id}`,
        source: getLastNode.id.toString(),
        target: (newNodeIds[0].id),
      }];

      setEdges((prevEdges) => [...prevEdges, ...newEdges]);

    }
    
  };

  const handleStop = async () => {
    if (initialNodes.length === 0 || actualState ==="Terminated" || actualState ==="Ready") return
    try {
      const response = await fetch(`http://localhost:8080/stopProcess?pid=${getPid}`)
      const data = await response.json()
      console.log(data.message)
    } catch (error) {
      console.log(error)
    } finally {
      setActualState('Ready')
      let positionX  = getLastNodePosition()[0]+200
      let positionY  = getLastNodePosition()[1]
      if ( positionX > 1000 ) {
        positionX = 0
        positionY = positionY + 100
      }

      const newNodeIds = ['Ready'].map((label) => {
        const id = getRandomInt(0,500).toString();
        return {
          id,
          position: { x: positionX , y: positionY }, // Initial position
          data: { label },
          style: { backgroundColor: '#f5f295'}
        };
      });

      setInitialNodes((prevNodes) => {
        // Update style of previous nodes
        const updatedNodes = prevNodes.map((node) => ({
          ...node,
          style: { backgroundColor: 'white' } // Change to white color
        }));
      
        // Append new nodes to the updated nodes
        return [...updatedNodes, ...newNodeIds];
      });

      const getLastNode = initialNodes[initialNodes.length - 1]
      const newEdges = [{
        id: `e${getLastNode.id}-${newNodeIds[0].id}`,
        source: getLastNode.id.toString(),
        target: (newNodeIds[0].id),
      }];

      setEdges((prevEdges) => [...prevEdges, ...newEdges]);

    }
    
  }

  const handleResume = async () => {
    if (initialNodes.length === 0 || actualState ==="Terminated" || actualState==="Running") return
    try {
      const response = await fetch(`http://localhost:8080/resumeProcess?pid=${getPid}`)
      const data = await response.json()
      console.log(data.message)
    } catch (error) {
      console.log(error)
    } finally {
      setActualState('Running')
      let positionX  = getLastNodePosition()[0]+200
      let positionY  = getLastNodePosition()[1]
      if ( positionX > 1000 ) {
        positionX = 0
        positionY = positionY + 100
      }

      const newNodeIds = ['Running'].map((label) => {
        const id = getRandomInt(0,500).toString();
        return {
          id,
          position: { x: positionX , y: positionY }, // Initial position
          data: { label },
          style: { backgroundColor: '#95f5a4' }
        };
      });

      setInitialNodes((prevNodes) => {
        // Update style of previous nodes
        const updatedNodes = prevNodes.map((node) => ({
          ...node,
          style: { backgroundColor: 'white' } // Change to white color
        }));
      
        // Append new nodes to the updated nodes
        return [...updatedNodes, ...newNodeIds];
      });

      const getLastNode = initialNodes[initialNodes.length - 1]
      const newEdges = [{
        id: `e${getLastNode.id}-${newNodeIds[0].id}`,
        source: getLastNode.id.toString(),
        target: (newNodeIds[0].id),
      }];

      setEdges((prevEdges) => [...prevEdges, ...newEdges]);
    
    }
  
  }

  return (
    <section className=' max-h-max '>
      <div className='flex justify-center gap-x-20 border p-5 '>
          <p className='text-xl p-4'> {
            getPid ? `PID: ${getPid}` : `PID: - - - -`
          }</p>
          <button onClick={handleNew} className='flex gap-x-4 p-4 cursor-pointer'>
            <BiSolidMessageSquareAdd size={20} />
            <p>New</p>
          </button>
          <button onClick={handleStop} className='flex gap-x-4 p-4 cursor-pointer'>
            <FaStopCircle size={20} />
            <p>Stop</p>
          </button>
          <button onClick={handleKill} className='flex gap-x-4 p-4 cursor-pointer'>
            <FaSkull size={20} />
            <p>Kill</p>
          </button>
          <button onClick={handleResume} className='flex gap-x-4 p-4 cursor-pointer'>
            <FaPlay size={20} />
            <p>Resume</p>
          </button>
          {actualState && <p className=' p-4'>Current State: {actualState}</p>}
      </div>
      <div className="flex p-5">
      <div className="w-1/5 text-justify ">
       <div className="border p-3 rounded-xl bg-sky-100">
        <p className="text-xl font-bold text-slate-600 text-center">Help Center</p>
        <p className=" text-slate-500 pt-2 ">Menu Options: </p>
        <p className=" text-slate-500 pt-2 "><strong>New:</strong> will create a new process,
        this will return the PID of the Process and the following states "New, Ready & Running"</p>
        <p className=" text-slate-500 pt-2 "><strong>Kill:</strong> Completly ends the process
        and return the state of "Terminated"</p>
        <p className=" text-slate-500 pt-2 "><strong>Stop:</strong> Changes the state from Running to Ready</p>
        <p className=" text-slate-500 pt-2 "><strong>Resume:</strong> Changes the state from Ready to Running</p>
       </div>

      </div>
      <div className="w-4/5 p-5 border ">
        <StateCanva initialNodes={initialNodes} initialEdges={initialEdges} />
      </div>
      </div>
    </section>
  )
}

export default SManagment