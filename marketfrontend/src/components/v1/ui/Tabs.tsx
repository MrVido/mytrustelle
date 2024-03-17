import React, { useState } from 'react';

interface Tab {
  id: string;
  title: string;
  content: React.ReactNode;
}

interface TabsProps {
  tabs: Tab[];
}

const Tabs: React.FC<TabsProps> = ({ tabs }) => {
  const [activeTabId, setActiveTabId] = useState(tabs[0].id);

  return (
    <div className="tabs">
      <div className="tab-titles">
        {tabs.map((tab) => (
          <button
            key={tab.id}
            className={tab.id === activeTabId ? 'active' : ''}
            onClick={() => setActiveTabId(tab.id)}
          >
            {tab.title}
          </button>
        ))}
      </div>
      <div className="tab-content">
        {tabs.find((tab) => tab.id === activeTabId)?.content}
      </div>
    </div>
  );
};

export default Tabs;
