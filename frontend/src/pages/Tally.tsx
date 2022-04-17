import React from 'react'
import { useTranslation } from 'react-i18next'

const Tally = () => {
  const { t } = useTranslation()
  
  return <div>
    {t('tallyPage')}
  </div>
} 

export default Tally