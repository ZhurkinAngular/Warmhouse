# SDK для сборки
FROM mcr.microsoft.com/dotnet/sdk:9.0-alpine AS build
WORKDIR /src

# Копируем всё и публикуем
COPY . .
RUN dotnet publish -c Release -o /app /p:UseAppHost=false

# Runtime-слой
FROM mcr.microsoft.com/dotnet/aspnet:9.0-alpine
WORKDIR /app

# Устанавливаем переменную окружения для прослушивания всех интерфейсов
ENV ASPNETCORE_URLS=http://+:8081

# Копируем опубликованное приложение
COPY --from=build /app .

# Документируем порт
EXPOSE 8081

# Запускаем
ENTRYPOINT ["dotnet", "temperature-api.dll"]
