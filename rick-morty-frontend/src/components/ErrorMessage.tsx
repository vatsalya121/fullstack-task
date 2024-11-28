import React from 'react';

interface ErrorMessageProps {
  message: string;
}

const ErrorMessage: React.FC<ErrorMessageProps> = ({ message }) => {
  return (
    <div className="bg-red-500 text-white p-2 rounded-md">
      <p>{message}</p>
    </div>
  );
};

export default ErrorMessage;
