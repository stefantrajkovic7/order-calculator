import './App.css'
import OrderCalculator from './OrderCalculator';
import PackSizesUpdater from './PackSizesUpdater';

function App() {
  return (
    <div className="min-h-screen bg-gray-100 flex flex-col items-center p-8">
      <div className="w-full max-w-xl">
        <PackSizesUpdater />
        <OrderCalculator />
      </div>
    </div>
  );
}

export default App
