import React, { useState, useEffect } from 'react';

const VoiceCommandListener = ({ onCommand }) => {
  const [listening, setListening] = useState(false);
  const [transcript, setTranscript] = useState('');

  useEffect(() => {
    // Check for browser support
    const SpeechRecognition = window.SpeechRecognition || window.webkitSpeechRecognition;
    if (!SpeechRecognition) {
      console.error('Speech Recognition API not supported in this browser.');
      return;
    }

    const recognition = new SpeechRecognition();
    recognition.continuous = true; // Keep listening even after the user stops speaking
    recognition.interimResults = true; // Show results that are not yet final

    // Event handler for when the recognition service returns a result
    recognition.onresult = (event) => {
      const transcriptArray = Array.from(event.results)
        .map(result => result[0])
        .map(result => result.transcript)
        .join('');
      setTranscript(transcriptArray);
      onCommand(transcriptArray); // Execute a callback function with the transcript
    };

    recognition.onerror = (event) => {
      console.error('Speech recognition error', event.error);
    };

    // Start or stop the speech recognition
    if (listening) {
      recognition.start();
    } else {
      recognition.stop();
    }

    return () => recognition.stop(); // Cleanup on unmount
  }, [listening, onCommand]);

  return (
    <div>
      <button onClick={() => setListening(prevState => !prevState)}>
        {listening ? 'Stop Listening' : 'Start Listening'}
      </button>
      <p>{transcript}</p>
    </div>
  );
};

export default VoiceCommandListener;
