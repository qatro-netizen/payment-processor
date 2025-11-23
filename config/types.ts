// types.ts

export type PaymentMethod = 'credit-card' | 'paypal' | 'bank-transfer';

export type PaymentStatus = 'pending' | 'confirmed' | 'rejected';

export type TransactionDetails = {
  id: string;
  paymentMethod: PaymentMethod;
  amount: number;
  status: PaymentStatus;
  date: string;
  paymentIntentId?: string;
};

export type PaymentIntent = {
  id: string;
  transactionId: string;
  status: PaymentStatus;
  paymentMethod: PaymentMethod;
  amount: number;
  date: string;
};

export type Customer = {
  id: string;
  name: string;
  email: string;
  address: string;
  phone: string;
};

export type PaymentProcessorConfig = {
  apiKey: string;
  apiSecret: string;
  paymentIntentTimeout: number;
  batchTimeout: number;
};

export type PaymentProcessorError = {
  code: number;
  message: string;
  transactionId?: string;
  paymentIntentId?: string;
};