import React from 'react'
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import Page404 from './pages/404'
import Home from './pages/Home'
import Tally from './pages/Tally'


function App() {
  return <div>
    <BrowserRouter>
      <Routes>
        <Route path="/tally" element={<Tally />} />
        <Route path="/" element={<Home />} />
        <Route path="*" element={<Page404 />} />
      </Routes>
    </BrowserRouter>
  </div>
}

export default App
