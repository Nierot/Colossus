// eslint-disable-next-line no-undef
module.exports = {
  content: [
    './src/**/*.{js,jsx,ts,tsx}',  
  ],
  theme: {
    extend: {
      boxShadow: {
        '2xl-white': '0 25px 50px 0 rgba(255, 255, 255, 0.25)',
      },
      fontFamily: {
        'player': ['Montserrat', 'ui-sans-serif', 'system-ui']
      }
    },
    colors: {
      'primary': '#59114D',
      'secondary': '#006C67',
      'accent': '#EE7B30',
      'secondary-accent': '#A7C957',
      'tertiary': '#73D2DE',
    }
  },
  plugins: [],
}
