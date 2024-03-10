import  { useEffect, useState,  } from 'react'
import "chart.js/auto"
import { Line } from 'react-chartjs-2';


const History = () => {
  const [getRamHistory, setRamHistory] = useState(null)


  useEffect(() => {
    fetchRamHistory()
  }, []);

  const fetchRamHistory = async () => {
    try {
      const response = await fetch('http://localhost:8080/getRam')
      const data = await response.json()
      const chartData = createData(data)
      setRamHistory(chartData)
    } catch (error) {
      console.log(error)
    }
  }

  function createData(rawJson) {
    const fechas = rawJson.data.map(d => d.fecha);
    const totalRamValues = rawJson.data.map(d => d.porcentaje);


    const data = {
      labels: fechas,
      datasets: [
        {
          label: 'Ram Consumed',
          fill: false,
          lineTension: 0.1,
          backgroundColor: 'rgba(75,192,192,0.4)',
          borderColor: 'rgba(75,192,192,1)',
          pointBorderColor: 'rgba(75,192,192,1)',
          pointBackgroundColor: 'rgba(255,255,255,0.4)',
          pointHoverRadius: 5,
          pointHoverBackgroundColor: 'rgba(255,99,132,1)',
          pointHoverBorderColor: 'rgba(255,99,132,1)',
          pointHitRadius: 10,
          data: totalRamValues
        }
      ]
    };


    return data;

  }

  

  return (
    <section className=' pt-12 pl-10 '>
      <h1 className='text-4xl  text-slate-700 font-bold mb-5'>History Ram & CPU</h1>
      <div className='flex '>
        <div className="w-1/2 border border-b-gray-600 shadow-xl rounded-xl  ">
          <h1 className='text-2xl text-center font-bold mb-5 pt-5'>Ram History</h1>
          
          <div className=' mt-4 mb-10 h-96 w-full flex items-center justify-center '>
            {
              getRamHistory && (
                <Line data={getRamHistory} />
              )
            }
           
          </div>
        </div>
        <div className="w-1/2">
          <p>hola</p>
        </div>
      </div>
    </section>
  )
}

export default History