import React from 'react';
import { Select } from '@grafana/ui';
import { MyQuery, JSONPlaceholderEntity } from '../types';

interface JSONPlaceholderEditorProps {
  query: MyQuery;
  onChange: (value: MyQuery) => void;
  onRunQuery: () => void;
}

const JsonPlaceHolderEntities = [
  { label: 'Todos', value: JSONPlaceholderEntity.Todos },
  { label: 'Users', value: JSONPlaceholderEntity.Users },
  { label: 'Posts', value: JSONPlaceholderEntity.Posts },
  { label: 'Comments', value: JSONPlaceholderEntity.Comments },
  { label: 'Albumns', value: JSONPlaceholderEntity.Albums },
  { label: 'Photos', value: JSONPlaceholderEntity.Photos },
];

export const JSONPlaceholderEditor: React.FC<JSONPlaceholderEditorProps> = props => {
  let { onChange, query, onRunQuery } = props;

  const onJSONPlaceholderEntityhange = (entity: JSONPlaceholderEntity) => {
    onChange({ ...query, jsonPlaceholderEntity: entity });
    onRunQuery();
  };

  return (
    <>
      <label className="gf-form-label query-keyword width-8">Type</label>
      <Select
        options={JsonPlaceHolderEntities}
        value={query.jsonPlaceholderEntity}
        onChange={e => onJSONPlaceholderEntityhange(e.value as JSONPlaceholderEntity)}
      ></Select>
    </>
  );
};
