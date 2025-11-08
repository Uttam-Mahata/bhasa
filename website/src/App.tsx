import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Layout from './components/Layout';
import Home from './pages/Home';
import Features from './pages/Features';
import Examples from './pages/Examples';
import GettingStarted from './pages/GettingStarted';
import Docs from './pages/Docs';

function App() {
  return (
    <Router>
      <Layout>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/features" element={<Features />} />
          <Route path="/examples" element={<Examples />} />
          <Route path="/getting-started" element={<GettingStarted />} />
          <Route path="/docs" element={<Docs />} />
        </Routes>
      </Layout>
    </Router>
  );
}

export default App;
