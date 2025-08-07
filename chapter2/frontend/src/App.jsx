import { useState, useEffect } from 'react';
import './App.css';

function SearchBar({ onSearch }) {
  const [searchTerm, setSearchTerm] = useState('中国');

  const handleSubmit = (e) => {
    e.preventDefault();
    onSearch(searchTerm);
  };

  return (
    <form className="search-bar" onSubmit={handleSubmit}>
      <input
        type="text"
        value={searchTerm}
        onChange={(e) => setSearchTerm(e.target.value)}
        placeholder="输入搜索关键词..."
      />
      <button type="submit">搜索</button>
    </form>
  );
}

function SearchResults({ results }) {
  if (results.length === 0) {
    return <div className="no-results">没有找到相关结果</div>;
  }

  return (
    <div className="results-container">
      <h2>搜索结果 ({results.length})</h2>
      <div className="results-list">
        {results.map((result) => (
          <div key={result.id} className="result-item">
            <h3>{result.title}</h3>
            <p className="description">{result.description}</p>
            <div className="metadata">
              <span className="source">{result.source}</span>
              <span className="date">{result.date}</span>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

function App() {
  const [results, setResults] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  // 实际API调用
  const search = async (term) => {
    setLoading(true);
    setError(null);

    try {
      // 调用后端API
      const response = await fetch(`http://localhost:8080/api/search?q=${encodeURIComponent(term)}`);
      if (!response.ok) {
        throw new Error(`搜索失败: ${response.status} ${response.statusText}`);
      }
      const data = await response.json();
      // 为结果添加id，以便React能正确渲染列表
      const resultsWithId = data.map((item, index) => ({
        ...item,
        id: index + 1
      }));
      setResults(resultsWithId);
    } catch (err) {
      setError(err.message || '搜索失败，请稍后再试');
      console.error('搜索错误:', err);
    } finally {
      setLoading(false);
    }
  };

  // 初始加载
  useEffect(() => {
    search('中国');
  }, []);

  return (
    <div className="app-container">
      <header>
        <h1>中文RSS搜索</h1>
      </header>
      <main>
        <SearchBar onSearch={search} />

        {loading && <div className="loading">加载中...</div>}
        {error && <div className="error">{error}</div>}
        {!loading && !error && <SearchResults results={results} />}
      </main>
      <footer>
        <p>© 2023 中文RSS搜索应用</p>
      </footer>
    </div>
  );
}

export default App
