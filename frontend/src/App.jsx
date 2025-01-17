import './App.css'
import OrderCalculator from './OrderCalculator';
import PackSizesUpdater from './PackSizesUpdater';

function App() {
  return (
    <div className="container mx-auto">
      <OrderCalculator />
      <PackSizesUpdater />
    </div>
  );
}

export default App
