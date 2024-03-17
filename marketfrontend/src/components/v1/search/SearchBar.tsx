import React, { useState } from 'react';

interface SearchBarProps {
  onSearch: (query: string) => void; // Callback function when a search is performed
}

const SearchBar: React.FC<SearchBarProps> = ({ onSearch }) => {
  const [query, setQuery] = useState('');

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    onSearch(query); // Call the onSearch prop with the current query state
  };

  return (
    <div className="search-bar">
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="Search..."
          value={query}
          onChange={(e) => setQuery(e.target.value)}
          aria-label="Search"
        />
        <button type="submit">Search</button>
      </form>
    </div>
  );
};

export default SearchBar;
