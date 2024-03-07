import React from 'react';

type AttributeValue = string | number | boolean | Date;
interface Attribute {
  label: string;
  value: AttributeValue;
  type: 'text' | 'number' | 'boolean' | 'date';
}

interface DynamicAttributesProps {
  attributes: Attribute[];
}

const formatAttributeValue = (value: AttributeValue, type: Attribute['type']): string => {
  switch (type) {
    case 'date':
      return value instanceof Date ? value.toLocaleDateString() : '';
    case 'boolean':
      return value === true ? 'Yes' : 'No';
    default:
      return value.toString();
  }
};

const DynamicAttributes: React.FC<DynamicAttributesProps> = ({ attributes }) => {
  return (
    <div className="dynamic-attributes">
      {attributes.map((attribute, index) => (
        <div key={index} className="attribute">
          <strong>{attribute.label}:</strong> {formatAttributeValue(attribute.value, attribute.type)}
        </div>
      ))}
    </div>
  );
};

export default DynamicAttributes;
