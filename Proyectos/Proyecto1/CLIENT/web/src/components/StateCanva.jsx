import React, { useEffect } from 'react'
import ReactFlow from 'reactflow';
 
import 'reactflow/dist/style.css';
const StateCanva = ({initialNodes, initialEdges}) => {

  useEffect(() => {
    console.log(initialNodes)
    console.log(initialEdges)
  }, [initialEdges])

  return (
    <div className='h-screen'>
        <ReactFlow nodes={initialNodes} edges={initialEdges} />
    </div>
  )
}

export default StateCanva