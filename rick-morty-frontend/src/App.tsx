import React, { useState, useEffect } from 'react';
import axios from 'axios';

type Character = {
  id: number;
  name: string;
  image: string;
  status: string;
  species: string;
  gender: string;
  episode: string[];
};

const App = () => {
  const [characters, setCharacters] = useState<Character[]>([]);
  const [searchQuery, setSearchQuery] = useState<string>('');
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string>('');

  // Fetch characters from your API (adjust the API URL as needed)
  useEffect(() => {
    setLoading(true);
    axios
      .get(`https://rickandmortyapi.com/api/character/?name=${searchQuery}`)
      .then((response) => {
        setCharacters(response.data.results);
        setLoading(false);
      })
      .catch((error) => {
        setError('Failed to fetch data');
        setLoading(false);
      });
  }, [searchQuery]);

  return (
    <div className="min-h-screen bg-gray-100 p-6">
      <div className="max-w-4xl mx-auto bg-white p-4 rounded-lg shadow-lg">
        <h1 className="text-3xl font-bold text-center text-gray-800 mb-4">
          Rick and Morty Character Search
        </h1>
        <div className="mb-6">
          <input
            type="text"
            className="w-full p-3 border border-gray-300 rounded-md"
            placeholder="Search characters..."
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
          />
        </div>

        {loading && <p className="text-center text-gray-600">Loading...</p>}
        {error && <p className="text-center text-red-500">{error}</p>}

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {characters.map((character) => (
            <div
              key={character.id}
              className="bg-white border border-gray-300 rounded-lg shadow-md p-4 hover:shadow-xl transition"
            >
              <img
                src={character.image}
                alt={character.name}
                className="w-full h-64 object-cover rounded-md mb-4"
              />
              <h2 className="text-xl font-semibold text-gray-800">{character.name}</h2>
              <p className="text-sm text-gray-600">Status: {character.status}</p>
              <p className="text-sm text-gray-600">Species: {character.species}</p>
              <p className="text-sm text-gray-600">Gender: {character.gender}</p>
              <div className="mt-2">
                <h3 className="text-sm font-semibold text-gray-700">Episodes:</h3>
                <ul className="list-disc list-inside">
                  {character.episode.slice(0, 5).map((ep, index) => (
                    <li key={index} className="text-sm text-gray-600">{ep}</li>
                  ))}
                </ul>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default App;
