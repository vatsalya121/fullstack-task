import React from 'react';
import CharacterCard from './CharacterCard';

interface CharacterListProps {
  characters: any[];
}

const CharacterList: React.FC<CharacterListProps> = ({ characters }) => {
  if (characters.length === 0) {
    return <p>No characters found.</p>;
  }

  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
      {characters.map((character) => (
        <CharacterCard key={character.id} character={character} />
      ))}
    </div>
  );
};

export default CharacterList;
