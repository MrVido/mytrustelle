import React from 'react';

interface NFTCertificateProps {
  tokenId: string;
  imageUrl: string;
  title: string;
  owner: string;
  blockchainExplorerUrl: string;
}

const NFTCertificate: React.FC<NFTCertificateProps> = ({
  tokenId,
  imageUrl,
  title,
  owner,
  blockchainExplorerUrl,
}) => {
  return (
    <div className="nft-certificate">
      <h2>{title}</h2>
      <img src={imageUrl} alt={title} />
      <p>Token ID: {tokenId}</p>
      <p>Owner: {owner}</p>
      <a href={blockchainExplorerUrl} target="_blank" rel="noopener noreferrer">
        View on Blockchain Explorer
      </a>
    </div>
  );
};

export default NFTCertificate;
