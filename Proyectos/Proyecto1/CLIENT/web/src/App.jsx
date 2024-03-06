import { useState } from "react"


function App() {

  const [getData, setData] = useState([])


  const fetchData = async () => {
   try {
    const response = await fetch('http://localhost:8080/data')
    const data = await response.json()
    setData(data.data)
   
   } catch (error) {
      setData("Error: " + " Please check the server is running or not."  )
   }
  }

  return (
    <>
      <button 
        className=" bg-green-700 h-10 w-20 m-2 rounded-full text-white"
        onClick={fetchData}
      >
        Api Call
      </button>
      <div>
        {
          getData.length >= 1 ?
           (
            getData.map((item, index) => {
              return (
                <div key={index} className="bg-gray-200 m-2 p-2 rounded-lg">
                  <h1 className="text-xl font-bold">{item.id}</h1>
                  <p className="text-sm">{item.name}</p>
                </div>
              )
            })
           ) : (
             <h1 className="text-2xl font-bold text-red-500">{getData}</h1>
           )
        }
      </div>
      
    </>
  )
}

export default App
