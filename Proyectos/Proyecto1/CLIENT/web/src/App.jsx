/*
import React, { useState, useEffect, useRef } from 'react';
import "chart.js/auto"
import { Doughnut } from 'react-chartjs-2';

const App = () => {
  // State variables for chart data
  const [getRamInfo, setRamInfo] = useState(null)
  const [getCpuInfo, setCpuInfo] = useState(null)
  const [chartData, setChartData] = useState({
    labels: ['Ram Libre', 'Ram Usuada'],
    datasets: [
      {
        label: 'Ram Monitor',
        data: [50, 50],
        backgroundColor: ['#647D87', '#647D87']
      },
    ],
  });
  const [chartDataCpu, setChartDataCpu] = useState({
    labels: ['Cpu Libre', 'Cpu Usuada'],
    datasets: [
      {
        label: 'Cpu Monitor',
        data: [50, 50],
        backgroundColor: ['#FA6607', '#FA6607']
      },
    ],
  });
  const chartRef = useRef(null); // Reference to the chart instance
  const chartRefCpu = useRef(null); // Reference to the chart instance

  /*function getRandomInt(min, max) {
    const minCeiled = Math.ceil(min);
    const maxFloored = Math.floor(max);
    return Math.floor(Math.random() * (maxFloored - minCeiled) + minCeiled); // The maximum is exclusive and the minimum is inclusive
  }*/
  
  /*
  // Function to fetch and update chart data
  const updateChartData = async () => {
    try {
      console.log("Fetching data from server")
      const response = await fetch('/insertRam')
      const data = await response.json()
      setRamInfo(data.message)
      const newData = {
        labels: [`Ram Libre`, 'Ram Usuada'],
        datasets: [
          {
            label: 'Ram Monitor',
            data: [data.message.libre, data.message.memoriaEnUso], // Replace with your new data values
            backgroundColor: ['#1D2B53', '#647D87'],
           
          },
        ],
      };

      setChartData(newData); 
    } catch (error) {
        console.log(error)
        console.log("Error: " + " Please check the server is running or not."  )
    } finally {
       // Update the chart instance using chartRef
      if (chartRef.current) {
        chartRef.current.update();
      }
    }
   
    
  };


  const updateChartDataCPU = async () => {
    try {
      console.log("Fetching data from server")
      const response = await fetch('/insertCpu')
      const data = await response.json()
      setCpuInfo(data.message)
      
      const newData = {
        labels: [`CPU Restante`, 'CPU Usuado'],
        datasets: [
          {
            label: 'Cpu Monitor',
            data: [100-data.message.cpu_porcentaje, data.message.cpu_porcentaje], // Replace with your new data values
            backgroundColor: ['#FA6607', '#647D87'],
           
          },
        ],
      };

      setChartDataCpu(newData); 
    } catch (error) {
        console.log(error)
        console.log("Error: " + " Please check the server is running or not."  )
    } finally {
       // Update the chart instance using chartRef
      if (chartRefCpu.current) {
        chartRefCpu.current.update();
      }
    }
   
    
  };
  

  // Run update function initially
  useEffect(() => {
    const interval = setInterval(() => {
      updateChartData();
      updateChartDataCPU();
    }, 3000)
    return () => clearInterval(interval)
  }, []);

  return (
    <section className=' pt-12 pl-10 '>
      <h1 className='text-4xl  text-slate-700 font-bold mb-5'>System Monitor</h1>
      <div className='flex '>
        <div className="w-1/2 border border-b-gray-600 shadow-xl rounded-xl  ">
          <h1 className='text-2xl text-center font-bold mb-5 pt-5'>RAM Monitor</h1>
          <div>
          
            {getRamInfo && (
              <div className='border rounded-xl p-5 m-5 bg-slate-50 flex flex-wrap justify-center '>
                <p className='p-2 text-lg'>Total Memory: {getRamInfo.totalRam}</p>
                <p className='p-2 text-lg'>Free Memory: {getRamInfo.libre}</p>
                <p className='p-2 text-lg'>Used Memory: {getRamInfo.memoriaEnUso}</p>
                <p className='p-2 text-lg'>Used: {getRamInfo.porcentaje} %</p>
              </div>
            )}
          </div>
          <div className=' mt-4 mb-10 h-60 '>
            <Doughnut updateMode='active' data={chartData} ref={chartRef} options={{ maintainAspectRatio: false }} />
          </div>
        </div>
        <div className="w-1/2 border border-b-gray-600 shadow-xl rounded-xl  ">
          <h1 className='text-2xl text-center font-bold mb-5 pt-5'>CPU Monitor</h1>
            <div className=' mt-4 mb-10 h-60 '>
             <Doughnut updateMode='active' data={chartDataCpu} ref={chartRef} options={{ maintainAspectRatio: false }} />
             </div>
        </div>
      </div>
    </section>
    
  );
};

export default App;

*/
import {useEffect, useState} from 'react';
import { CircularProgressbar } from 'react-circular-progressbar';

const App = () => {
  const [getRamInfo, setRamInfo] = useState(null)
  const [getCpuInfo, setCpuInfo] = useState(null)

  
  const updateChartData = async () => {
    try {
      console.log("Fetching data from server")
      const response = await fetch('/insertRam')
      const data = await response.json()
      setRamInfo(data.message)
    } catch (error) {
        console.log(error)
        console.log("Error: " + " Please check the server is running or not."  )
    }
  };


  const updateChartDataCPU = async () => {
    try {
      console.log("Fetching data from server")
      const response = await fetch('/insertCpu')
      const data = await response.json()
      setCpuInfo(data.message)
    } catch (error) {
        console.log(error)
        console.log("Error: " + " Please check the server is running or not."  )
    }
  }

  useEffect(() => {
    const interval = setInterval(() => {
      updateChartData();
      updateChartDataCPU();
    }, 500)
    return () => clearInterval(interval)
  }, []);


  return (

    <section className=' pt-12 pl-10 '>
      <h1 className='text-4xl  text-slate-700 font-bold mb-5'>System Monitor</h1>
      <div className='flex '>
        <div className="w-1/2 border border-b-gray-600 shadow-xl rounded-xl  ">
          <h1 className='text-2xl text-center font-bold mb-5 pt-5'>RAM Monitor</h1>
          <div className=' flex items-center justify-center mt-4 mb-10 h-60 '>
            {getRamInfo ? (
                <div className=' h-60 w-60'>
                <CircularProgressbar 
                  value={getRamInfo.porcentaje} 
                  text={`${getRamInfo.porcentaje}%`}
                />
              </div>
            ): (
              <div className=' h-60 w-60'>
                  <CircularProgressbar 
                  value={0} 
                  text={`0%`}
                />
              </div>
            )}
          </div>
        </div>
        <div className="w-1/2 border border-b-gray-600 shadow-xl rounded-xl  ">
          <h1 className='text-2xl text-center font-bold mb-5 pt-5'>CPU Monitor</h1>
            <div className=' flex items-center justify-center mt-4 mb-10 h-60 '>
             {getCpuInfo ? (
                <div className=' h-60 w-60'>
                    <CircularProgressbar 
                    value={getCpuInfo.cpu_porcentaje} 
                    text={`${getCpuInfo.cpu_porcentaje}%`}
                  />
                </div>
             ): (
              <div className=' h-60 w-60'>
                  <CircularProgressbar 
                  value={0} 
                  text={`0%`}
                />
              </div>
             
             )}
             </div>
        </div>
      </div>
    </section>

  )
};

export default App;
