# SDK для сборки
FROM mcr.microsoft.com/dotnet/sdk:9.0-alpine AS build
WORKDIR /src

# Копируем проект и устанавливаем тулзу
COPY . .
RUN dotnet new tool-manifest --force
RUN dotnet tool install Swashbuckle.AspNetCore.Cli --version 9.0.3 --local

# Сборка (именно Build, не Publish, т.к. Publish удаляет dev-зависимости)
RUN dotnet build -c Release

# Генерация Swagger JSON
RUN dotnet tool run swagger tofile \
    --output /src/swagger.json \
    ./bin/Release/net9.0/temperature-api.dll \
    v1

# Публикация после Swagger-генерации
RUN dotnet publish -c Release -o /app /p:UseAppHost=false

# Runtime-слой
FROM mcr.microsoft.com/dotnet/aspnet:9.0-alpine
WORKDIR /app

ENV ASPNETCORE_URLS=http://+:8081

COPY --from=build /app .
COPY --from=build /src/swagger.json ./swagger.json

EXPOSE 8081
ENTRYPOINT ["dotnet", "temperature-api.dll"]
