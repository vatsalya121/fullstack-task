import React from 'react';

interface SearchBarProps {
  searchTerm: string;
  onSearch: (term: string) => void;
}

const SearchBar: React.FC<SearchBarProps> = ({ searchTerm, onSearch }) => {
  return (
    <div className="mb-4">
      <input
        type="text"
        placeholder="Search for a character"
        value={searchTerm}
        onChange={(e) => onSearch(e.target.value)}
        className="p-2 border border-gray-300 rounded-md w-full"
      />
    </div>
  );
};

export default SearchBar;
