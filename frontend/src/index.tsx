import React from 'react'
import ReactDOM from 'react-dom'
import './index.css'
import App from './App'
import i18next from 'i18next'
import { initReactI18next } from 'react-i18next'

const translation = require('./translation.json')

i18next
  .use(initReactI18next)
  .init(translation)

ReactDOM.render(
  <App />,
  document.getElementById('root')
)