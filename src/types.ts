import { DataQuery, DataSourceJsonData } from '@grafana/data';

export enum EntitiyType {
  Dummy = 'dummy',
  Todos = 'todos',
}

export interface MyQuery extends DataQuery {
  entityType: EntitiyType;
  queryText?: string;
  constant?: number;
  numberOfTodos?: number;
  hideFinishedTodos?: boolean;
}

export const defaultQuery: Partial<MyQuery> = {
  entityType: EntitiyType.Dummy,
};

export interface MyDataSourceOptions extends DataSourceJsonData {
  path?: string;
}

export interface MySecureJsonData {
  apiKey?: string;
}
