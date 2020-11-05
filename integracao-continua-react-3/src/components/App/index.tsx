import React, { useState } from 'react';
import { Chart } from '../Chart';
import { Header } from '../Header';
import './index.css';

function App() {
  const [coinSelected, setCoinSelected] = useState("BTC");
  return (
    <div className="App">
      <Header onSelected={(coin) => setCoinSelected(coin)}/>
      <Chart coin={coinSelected}/>
    </div>
  );
}

export default App;
