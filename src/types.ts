import { DataQuery, DataSourceJsonData } from '@grafana/data';

export enum EntityType {
  Dummy = 'dummy',
  Todos = 'todos',
  JSONPlaceholder = 'jsonplaceholder',
}
export enum JSONPlaceholderEntity {
  Todos = 'todos',
  Users = 'users',
}

export interface MyQuery extends DataQuery {
  entityType: EntityType;
  queryText?: string;
  constant?: number;
  numberOfTodos?: number;
  hideFinishedTodos?: boolean;
  jsonPlaceholderEntity?: JSONPlaceholderEntity;
}

export const defaultQuery: Partial<MyQuery> = {
  entityType: EntityType.Dummy,
  numberOfTodos: 200,
  hideFinishedTodos: false,
  jsonPlaceholderEntity: JSONPlaceholderEntity.Todos,
};

export interface MyDataSourceOptions extends DataSourceJsonData {
  path?: string;
}

export interface MySecureJsonData {
  apiKey?: string;
}
