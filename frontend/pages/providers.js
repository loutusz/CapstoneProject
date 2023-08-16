// pages/index.js
import React from 'react';
import ProviderCard from '@/components/providercard';

const providers = [
  'Telecommunications',
  'Internet Service Providers',
  'Streaming Services',
  'Cloud Storage',
  'Payment Processors',
  'Web Hosting',
  'Health Insurance',
  'Car Insurance',
  'Online Marketplaces',
  'Travel Booking',
  // Add more providers here
];

const Providers = () => {
  return (
    <div className="flex justify-center items-center h-screen bg-gray-100">
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {providers.map(provider => (
          <ProviderCard key={provider} providerName={provider} />
        ))}
      </div>
    </div>
  );
};

export default Providers;
