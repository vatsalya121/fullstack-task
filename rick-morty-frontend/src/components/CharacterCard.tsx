import React from 'react';

interface CharacterCardProps {
  character: {
    id: number;
    name: string;
    image: string;
    episode: string[];
  };
}

const CharacterCard: React.FC<CharacterCardProps> = ({ character }) => {
  return (
    <div className="bg-white p-4 rounded-lg shadow-md">
      <img src={character.image} alt={character.name} className="w-full h-48 object-cover rounded-md mb-4" />
      <h2 className="text-xl font-semibold mb-2">{character.name}</h2>
      <p>Episodes: {character.episode.length}</p>
    </div>
  );
};

export default CharacterCard;
