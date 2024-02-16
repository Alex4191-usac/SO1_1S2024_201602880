import {useEffect, useState} from 'react';
import {ReadRam} from "../wailsjs/go/main/App";
import { CircularProgressbar } from 'react-circular-progressbar';
import './App.css';

function App() {
    const [resultRam, setResultRam] = useState(null);
    const [loaded, setLoaded] = useState(false);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const result = await ReadRam();
                setResultRam(result);
            } finally {
                setLoaded(true);
            }
        };
        const interval = setInterval(() => {
            fetchData();
        }, 500);

        return () => clearInterval(interval);
    }, []);

   
    return (
        <div id="App">
            <h1>Wails React App</h1>
            <div className="ram">
                <h2>RAM</h2>
                {resultRam === null ? <p>Loading...</p> :
                    <div id="resultRam" className="result">
                        <p>Free Memory : {resultRam.libre} </p>
                        <p>UsedMemory: {resultRam.memoriaEnUso}</p>
                        {loaded ?
                            (
                                <div className="graph-pie-container">
                                    
                                    <div className="graph-pie">
                                        <CircularProgressbar 
                                            value={resultRam.porcentaje} 
                                            text={`${resultRam.porcentaje}%`}
                                        />
                                    </div>
                                   
                                </div>
                            ) : <p>Loading...</p>}
                    </div>

                 
                }
            </div>
        </div>
    );
}

export default App;
