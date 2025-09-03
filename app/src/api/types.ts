// API Types and Interfaces

export interface Trip {
  id: number;
  user_id: number;
  origin: string;
  destination: string;
  departure_time: number; // Unix timestamp
  arrival_time: number;
  airline: string;
  flight_number: string;
  reservation?: string;
  terminal?: string;
  gate?: string;
  origin_name?: string;
  destination_name?: string;
  origin_lat?: number;
  origin_lng?: number;
  destination_lat?: number;
  destination_lng?: number;
}

export interface ConnectingTrip {
  from_trip: Trip;
  to_trip: Trip;
}

export interface ApiResponse<T> {
  success: boolean;
  data?: T;
  error?: {
    code: string;
    message: string;
  };
  meta?: {
    page?: number;
    limit?: number;
    total?: number;
  };
}

export interface AuthTokens {
  access_token: string;
  refresh_token: string;
  expires_in: number;
}

export interface User {
  id: number;
  name: string;
  email: string;
  picture?: string;
}

export interface Statistics {
  total_trips: number;
  countries_visited: number;
  total_miles: number;
  favorite_airline: string;
  upcoming_trips: number;
  past_trips: number;
}

export interface Airport {
  iata_code: string;
  name: string;
  city: string;
  country: string;
  latitude: number;
  longitude: number;
}

export interface FlightData {
  flight_number: string;
  airline: string;
  origin: string;
  destination: string;
  departure_time: string;
  arrival_time: string;
  status: string;
}

// Request/Response types
export interface CreateTripRequest {
  origin: string;
  destination: string;
  departure_time: number;
  arrival_time: number;
  airline: string;
  flight_number: string;
  reservation?: string;
  terminal?: string;
  gate?: string;
}

export interface UpdateTripRequest extends Partial<CreateTripRequest> {
  id: number;
}

export interface LoginRequest {
  google_token: string;
}

export interface LoginResponse {
  success: boolean;
  data?: {
    user: User;
    tokens: AuthTokens;
  };
  error?: {
    code: string;
    message: string;
  };
}